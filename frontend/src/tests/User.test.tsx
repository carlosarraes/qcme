import { screen } from '@testing-library/react'
import renderWithRouter from '../mocks/renderWithRouter'
import User from '../pages/User'

describe('User Page tests', () => {
  beforeEach(() => {
    renderWithRouter(<User />, { route: '/john', path: '/:name' })
  })

  test('renders loading correctly', () => {
    const loading = screen.getByText('Loading...')

    expect(loading).toBeInTheDocument()
  })

  test('renders user page correctly', async () => {
    const name = await screen.findByText(/John/i)
    const gitHub = await screen.findByText(/github/i)
    const linkedIn = await screen.findByText(/linkedin/i)

    expect(name).toBeInTheDocument()
    expect(gitHub).toBeInTheDocument()
    expect(linkedIn).toBeInTheDocument()
  })

  test('renders user not found message', async () => {
    renderWithRouter(<User />, { route: '/unknown', path: '/:name' })

    const userNotFound = await screen.findByTestId('user-not-found')

    expect(userNotFound).toBeInTheDocument()
  })
})
