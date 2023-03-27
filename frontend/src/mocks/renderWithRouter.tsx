/* c8 ignore start */
import { ReactNode } from 'react'
import { render, RenderResult } from '@testing-library/react'
import { BrowserRouter, useNavigate, useLocation, Route, Routes } from 'react-router-dom'

interface RenderWithRouterOptions {
  route?: string
  path?: string
}

interface RenderWithRouterResult extends RenderResult {
  navigate: ReturnType<typeof useNavigate>
  location: ReturnType<typeof useLocation>
}

function renderWithRouter(
  ui: ReactNode,
  { route = '/', path }: RenderWithRouterOptions = {},
): RenderWithRouterResult {
  window.history.pushState({}, 'Test page', route)

  let navigate: ReturnType<typeof useNavigate> | null = null
  let location: ReturnType<typeof useLocation> | null = null

  const Wrapper = ({ children }: { children: ReactNode }) => {
    navigate = useNavigate()
    location = useLocation()

    return <>{children}</>
  }

  const result = render(
    <BrowserRouter>
      <Wrapper>
        <Routes>
          <Route path={path} element={ui} />
        </Routes>
      </Wrapper>
    </BrowserRouter>,
  )

  if (!navigate || !location) {
    throw new Error('Navigate and location should be defined')
  }

  return { ...result, navigate, location }
}

export default renderWithRouter
/* c8 ignore end */
