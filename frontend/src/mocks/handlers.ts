/* c8 ignore start */
import { rest } from 'msw'

const url = 'https://goqr-4wgfen3n5q-rj.a.run.app/'

const mockUser = {
  name: 'John',
  linkedIn: 'https://www.linkedin.com/in/johndoe/',
  gitHub: 'https://github.com/johndoe',
}

export const handlers = [
  rest.post(url + 'qrcodeme', (req, res, ctx) => {
    const { name } = req.body as { name: string }

    if (name === 'existingUser') {
      return res(ctx.status(400), ctx.json({ error: 'User already exists!' }))
    }

    const fakeQrCodeBlob = new Blob(['fakeQrCode'], { type: 'image/png' })
    return res(ctx.status(200), ctx.body(fakeQrCodeBlob))
  }),
  rest.get(url + ':name', (req, res, ctx) => {
    console.log(req.params)
    const { name } = req.params

    if (name === 'john') {
      return res(ctx.status(200), ctx.json(mockUser))
    } else {
      return res(ctx.status(404), ctx.json({ error: 'User not found!' }))
    }
  }),
]
/* c8 ignore end */
