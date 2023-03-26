import { useParams } from "react-router-dom"

const User = () => {
  const { name } = useParams<{ name: string }>()
  
  console.log(name)
  return <div>User</div>
}

export default User
