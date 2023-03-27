import { render, screen, waitFor } from '@testing-library/react'
import userEvent from '@testing-library/user-event'
import { BrowserRouter } from 'react-router-dom'
import { vi } from 'vitest'
import Home from '../pages/Home'

const renderWithRouter = (component: JSX.Element) => {
  return render(<BrowserRouter>{component}</BrowserRouter>)
}

const readAsDataURL = vi.fn()
vi.spyOn(FileReader.prototype, 'readAsDataURL').mockImplementation(readAsDataURL)

describe('Home', () => {
  beforeEach(() => {
    renderWithRouter(<Home />)
  })

  test('renders form correctly', () => {
    const nameLabel = screen.getByText('Name')
    const gitHubLabel = screen.getByText('GitHub')
    const linkedInLabel = screen.getByText('LinkedIn')
    const submitButton = screen.getByRole('button', { name: /submit/i })

    expect(nameLabel).toBeInTheDocument()
    expect(gitHubLabel).toBeInTheDocument()
    expect(linkedInLabel).toBeInTheDocument()
    expect(submitButton).toBeInTheDocument()
    expect(submitButton).toBeDisabled()
  })

  test('if button enables when all fields are filled', async () => {
    const submitButton = screen.getByRole('button', { name: /submit/i })
    const nameInput = screen.getByPlaceholderText('Used for your link')
    const gitHubInput = screen.getByPlaceholderText('Your GitHub')
    const linkedInInput = screen.getByPlaceholderText('Your LinkedIn')

    await userEvent.type(nameInput, 'test')
    expect(submitButton).toBeDisabled()

    await userEvent.type(gitHubInput, 'test')
    expect(submitButton).toBeDisabled()

    await userEvent.type(linkedInInput, 'test')
    expect(submitButton).toBeEnabled()
  })

  test('if button disables when one field is empty', async () => {
    const submitButton = screen.getByRole('button', { name: /submit/i })
    const nameInput = screen.getByPlaceholderText('Used for your link')
    const gitHubInput = screen.getByPlaceholderText('Your GitHub')
    const linkedInInput = screen.getByPlaceholderText('Your LinkedIn')

    await userEvent.type(nameInput, 'test')
    await userEvent.type(gitHubInput, 'test')
    await userEvent.type(linkedInInput, 'test')
    expect(submitButton).toBeEnabled()

    await userEvent.clear(nameInput)
    expect(submitButton).toBeDisabled()

    await userEvent.clear(gitHubInput)
    expect(submitButton).toBeDisabled()

    await userEvent.clear(linkedInInput)
    expect(submitButton).toBeDisabled()
  })

  test('if forms submission fails if user already exists', async () => {
    const submitButton = screen.getByRole('button', { name: /submit/i })
    const nameInput = screen.getByPlaceholderText('Used for your link')
    const gitHubInput = screen.getByPlaceholderText('Your GitHub')
    const linkedInInput = screen.getByPlaceholderText('Your LinkedIn')

    await userEvent.type(nameInput, 'existingUser')
    await userEvent.type(gitHubInput, 'test')
    await userEvent.type(linkedInInput, 'test')
    expect(submitButton).toBeEnabled()

    await userEvent.click(submitButton)

    const toast = await screen.findByText('User already exists!')
    expect(toast).toBeInTheDocument()
  })

  test('if forms submission succeeds if user doesnt exist', async () => {
    const submitButton = screen.getByRole('button', { name: /submit/i })
    const nameInput = screen.getByPlaceholderText('Used for your link')
    const gitHubInput = screen.getByPlaceholderText('Your GitHub')
    const linkedInInput = screen.getByPlaceholderText('Your LinkedIn')

    await userEvent.type(nameInput, 'NewUser')
    await userEvent.type(gitHubInput, 'test')
    await userEvent.type(linkedInInput, 'test')
    expect(submitButton).toBeEnabled()

    await userEvent.click(submitButton)

    waitFor(async () => {
      const downloadLink = await screen.findByText(/Download QR Code/i)
      expect(downloadLink).toBeInTheDocument()
    })
  })
})
