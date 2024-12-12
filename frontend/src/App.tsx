import { BrowserRouter, Route, Routes } from "react-router-dom"
import Landing from "./pages/Landing"
import Home from "./pages/Home"
import Auth from "./pages/Auth"



function App() {

  return (
    <>
    <BrowserRouter>
    
        <Routes>
          <Route path="/" element={<Landing/>}/>
          <Route path="/home" element={<Home/>}/>
          <Route path="/auth" element={<Auth/>}/>

        </Routes>
    
    </BrowserRouter>

    </>
  )
}

export default App
