import { useParams } from 'react-router-dom'
import { AiFillLinkedin, AiFillGithub } from 'react-icons/ai'
import { useEffect, useState } from 'react'
import { URL, UserType } from './Home'
import Spinner from '../components/Spinner'

const User = () => {
  const [user, setUser] = useState<UserType | null>(null)
  const [notFound, setNotFound] = useState<boolean>(false)
  const [loading, setLoading] = useState<boolean>(true)
  const { name } = useParams<{ name: string }>()

  useEffect(() => {
    const fetchData = async () => {
      const lowerCaseName = name!.toLowerCase()
      try {
        const response = await fetch(URL + lowerCaseName)
        if (!response.ok) {
          throw new Error('User not found')
        }
        const data = await response.json()
        setUser(data)
        setLoading(false)
      } catch (e) {
        setLoading(false)
        setNotFound(true)
      }
    }

    fetchData()
  }, [])

  if (loading) {
    return <Spinner />
  }

  if (notFound) {
    return (
      <main className="flex h-full w-11/12 max-w-sm mx-auto justify-center items-center">
        <section className="w-full flex flex-col items-start p-4 border-2 border-slate-400 rounded-md shadow-lg">
          <h2 className="text-2xl">
            User <span className="font-bold capitalize">{name}</span> was not found :(
          </h2>
        </section>
      </main>
    )
  }

  return (
    <main className="flex h-full w-11/12 max-w-sm mx-auto justify-center items-center">
      <section className="w-full flex flex-col items-start p-4 border-2 border-slate-400 rounded-md shadow-lg">
        <h2 className="text-2xl">Hello! :)</h2>
        <p className="text-md mt-6">
          Here are the links to <span className="font-bold capitalize">{user?.name}</span> social
          media accounts:
        </p>
        <div className="flex justify-between w-full mt-6">
          <a
            href={user?.linkedIn}
            className="flex w-40 h-10 items-center justify-center gap-2 bg-[#0072b1] text-white rounded-md shadow-md opacity-80 hover:opacity-100 duration-200 cursor-pointer"
          >
            <AiFillLinkedin className="text-2xl" />
            <span>LinkedIn</span>
          </a>
          <a
            href={user?.gitHub}
            className="flex w-40 h-10 items-center justify-center gap-2 bg-[#171515] text-white rounded-md shadow-md opacity-80 hover:opacity-100 duration-200 cursor-pointer"
          >
            <AiFillGithub className="text-2xl" />
            <span>GitHub</span>
          </a>
        </div>
      </section>
    </main>
  )
}

export default User
