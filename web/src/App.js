import './App.css';

import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import About from './pages/About';
import Login from './pages/Login';
import Navbar from './components/Navbar'
import Admin_Login from './pages/Admin_Login'
import Admin_Dashboard from './pages/Admin_Dashboard'
import Manage_Students from './pages/Admin_Pages/Manage_Student';

function App() {
  return (
    <Router>
      <>
        <header>
          <Navbar />
        </header>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About />} />
          <Route path="/login" element={<Login />} />
          <Route path='/admin/login' element={<Admin_Login />} />
          <Route path='/admin/dashboard' element={<Admin_Dashboard />} />
          <Route path='/manage/students' element={<Manage_Students />} />
        </Routes>
      </>
    </Router>

  );
}

export default App;
