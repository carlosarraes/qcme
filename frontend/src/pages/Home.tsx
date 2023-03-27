import { useEffect, useState } from 'react'
import { ToastContainer, toast } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'

export const URL = 'https://goqr-4wgfen3n5q-rj.a.run.app/'

export type UserType = {
  name: string
  gitHub: string
  linkedIn: string
}

const Home = () => {
  const [imageData, setImageData] = useState<string>('')
  const [user, setUser] = useState<UserType>({
    name: '',
    gitHub: '',
    linkedIn: '',
  })
  const [btnActive, setBtnActive] = useState<boolean>(false)

  useEffect(() => {
    validateBtn()
  }, [user])

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target
    setUser((prev) => ({ ...prev, [name]: value.trim() }))
  }

  const validateBtn = () => {
    const { name, gitHub, linkedIn } = user
    if (name && gitHub && linkedIn) {
      setBtnActive(true)
    } else {
      setBtnActive(false)
    }
  }

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    try {
      const response = await fetch(URL + 'qrcodeme', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(user),
      })
      if (!response.ok) {
        throw new Error()
      }

      const blob = await response.blob()
      const reader = new FileReader()
      reader.onloadend = () => {
        setImageData(reader.result as string)
      }
      reader.readAsDataURL(blob)
    } catch (e) {
      toast.error('User already exists!')
    }
  }

  const { name, gitHub, linkedIn } = user

  if (imageData) {
    toast.success('QR Code generated successfully!')
    return (
      <main className="flex h-full w-11/12 max-w-sm mx-auto justify-center items-center">
        <section className="w-full flex flex-col items-start p-4 border-2 border-slate-400 rounded-md shadow-lg">
          <>
            <div className="flex flex-col items-center mt-4 mx-auto">
              <img src={imageData} alt="QRCode" />
            </div>
            <a
              href={imageData}
              download="qrcode.png"
              className="mt-4 btn btn-outline btn-success w-full duration-200 shadow-md"
            >
              Download QR Code
            </a>
          </>
        </section>
        <ToastContainer position="top-center" />
      </main>
    )
  }

  return (
    <main className="flex h-full w-11/12 max-w-sm mx-auto justify-center items-center">
      <section className="w-full flex flex-col items-start p-4 border-2 border-slate-400 rounded-md shadow-lg">
        <h1 className="text-4xl font-bold mx-auto">QrCode Me</h1>
        <hr className="border-2 w-full border-slate-300 shadow-md mb-6" />
        <form onSubmit={handleSubmit} className="flex flex-col w-full">
          <div className="form-control gap-2">
            <label className="input-group">
              <span>Name</span>
              <input
                type="text"
                placeholder="Used for your link"
                className="input input-bordered w-full !outline-none"
                name="name"
                value={name}
                onChange={handleChange}
              />
            </label>
            <label className="input-group">
              <span>GitHub</span>
              <input
                type="text"
                placeholder="Your GitHub"
                className="input input-bordered w-full !outline-none"
                name="gitHub"
                value={gitHub}
                onChange={handleChange}
              />
            </label>
            <label className="input-group">
              <span>LinkedIn</span>
              <input
                type="text"
                placeholder="Your LinkedIn"
                className="input input-bordered w-full !outline-none"
                name="linkedIn"
                value={linkedIn}
                onChange={handleChange}
              />
            </label>
            <button
              type="submit"
              className="mt-4 btn btn-outline btn-ghost duration-200 shadow-md"
              disabled={!btnActive}
            >
              Submit
            </button>
          </div>
        </form>
      </section>
      <ToastContainer position="top-center" />
    </main>
  )
}

export default Home
