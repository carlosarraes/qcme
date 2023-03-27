import { rest } from 'msw'

const url = 'https://goqr-4wgfen3n5q-rj.a.run.app/'

export const handlers = [
  rest.post(url + 'qrcodeme', (req, res, ctx) => {
    const { name } = req.body as { name: string }

    if (name === 'existingUser') {
      return res(ctx.status(400), ctx.json({ error: 'User already exists!' }))
    }

    const fakeQrCodeBlob = new Blob(['fakeQrCode'], { type: 'image/png' })
    return res(ctx.status(200), ctx.body(fakeQrCodeBlob))
  }),
]
