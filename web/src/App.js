import './App.css';

import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import About from './pages/About';
import Login from './pages/Login';
import Navbar from './components/Navbar';

import Admin_Login from './pages/Admin_Login';
import Admin_Dashboard from './pages/Admin_Dashboard';

import Manage_Students from './pages/Admin_Pages/Manage_Students';
import Edit_Student from './pages/Admin_Pages/Edit_Student';

import Manage_Librarians from './pages/Admin_Pages/Manage_Librarians';
import Edit_Librarian from './pages/Admin_Pages/Edit_Librarian';

import Librarian_Manage_Books from './pages/Librarian_Pages/Manage_Books';
import Librarian_Edit_Books from './pages/Librarian_Pages/Edit_Books'
import Librarian_Manage_Students from './pages/Librarian_Pages/Manage_Students';
import Librarian_Edit_Student from './pages/Librarian_Pages/Edit_Student';

import Explore_Books from './pages/Student_Pages/Explore_Books.jsx'

import { AuthProvider } from './AuthContext.js';

function App() {
  return (
    <AuthProvider>
      <Router>
        <header>
          <Navbar />
        </header>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About />} />
          <Route path="/login" element={<Login />} />

          <Route path="/admin/login" element={<Admin_Login />} />
          <Route path="/admin/dashboard" element={<Admin_Dashboard />} />

          <Route path="/manage/students" element={<Manage_Students />} />
          <Route path="/edit/student/:id" element={<Edit_Student />} />

          <Route path="/manage/librarians" element={<Manage_Librarians />} />
          <Route path="/edit/librarian/:id" element={<Edit_Librarian />} />

          <Route path="/librarian/manage/books" element={<Librarian_Manage_Books />} />
          <Route path="/librarian/edit/book/:id" element={<Librarian_Edit_Books />} />

          <Route path="/librarian/manage/students" element={<Librarian_Manage_Students />} />
          <Route path="/librarian/edit/student/:id" element={<Librarian_Edit_Student />} />

          <Route path="/student/books" element={<Explore_Books />} />
        </Routes>
      </Router>
    </AuthProvider>
  );
}

export default App;
