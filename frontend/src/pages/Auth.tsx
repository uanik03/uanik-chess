import React, { useState } from 'react'

const Auth = () => {
    const [name, setName] = useState("")
    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [isLogin, setIsLogin] = useState(true)

    const handleSubmit = async()=>{
        if(isLogin){

        }else{

        }
    }


  return (
    <div className='flex items-center justify-center h-screen w-full'>
        <form onSubmit={handleSubmit} className='flex flex-col gap-5 w-1/4 h-3/5  rounded-lg items-center justify-start pt-8' style={{"backgroundColor":"rgb(38,38,38"}}>
        <h2 className='text-4xl'>Uanik-Chess</h2>
        <h3 className='text-3xl'>{isLogin?"Login":"Sign Up"}</h3>
            {!isLogin && <label htmlFor="name" className='w-full flex items-center justify-center'>
                <input type="text" name="name" value={name} placeholder='name' onChange={(e)=>setName(e.target.value)} className='p-2 h-11 rounded-md mb-4 w-3/4 bg-gray-200 ' style={{"color":"rgb(38,38,38"}} />
            </label>}
            <label htmlFor="email" className='w-full flex items-center justify-center'>
                <input type="email" name="email" value={email} placeholder='email' onChange={(e)=>setEmail(e.target.value)}  className='p-2 h-11 rounded-md mb-4 w-3/4 bg-gray-200' style={{"color":"rgb(38,38,38"}} />
            </label> 
            <label htmlFor="password" className='w-full flex items-center justify-center'>
                <input type="password" name="password" value={password} placeholder='password' onChange={(e)=>setPassword(e.target.value)}  className='p-2 h-11 rounded-md mb-4 w-3/4 bg-gray-200' style={{"color":"rgb(38,38,38"}} />
            </label>
            <button type="submit" className='bg-gray-200 hover:bg-gray-400 w-24 h-9 rounded-md text-black'>{isLogin? "login":"signUp"}</button>
        <p>{isLogin ?"Don't have an account ? ":"Already have an account? "} <span className='cursor-pointer' onClick={()=>{setIsLogin(!isLogin)}}>{isLogin ?" signup":" login"}</span></p>
        </form>

    </div>
  )
}

export default Auth