import { Link, useNavigate } from 'react-router-dom';
import { useAuth } from "../AuthContext";

export default function Navbar() {

  const { isAdminLogged, isStudentLogged, isLibrarianLogged } = useAuth()

  const navigate = useNavigate()

  const handleAdminLogout = async () => {
    await fetch(`/admin/logout`, {
      method: "POST",
    });
    navigate(`/admin/login`)
  };

  const handleUserLogout = async () => {
    await fetch(`/logout`, {
      method: "POST",
    });
    navigate(`/login`)
  };

  return (
    <nav>
      {isAdminLogged ? (
        <Link to="/admin/dashboard">Home</Link>
      ) : (
        <Link to="/">Home</Link>
      )}

      {!isAdminLogged && <Link to="/admin/login">Admin Login</Link>}

      {isAdminLogged && <Link to="/manage/students">Manage Students</Link>}
      {isAdminLogged && <Link to="/manage/librarians">Manage Librarian</Link>}

      {isAdminLogged && (
        <button onClick={handleAdminLogout} style={{ marginLeft: "10px", cursor: "pointer" }}>
          Logout Admin
        </button>
      )}

      <br />

      {!(isStudentLogged || isLibrarianLogged) && <Link to="/login">Login</Link>}

      {isLibrarianLogged && <Link to="/librarian/manage/students">Manage Students</Link>}
      {isLibrarianLogged && <Link to="/librarian/manage/books">Manage Books</Link>}

      {isStudentLogged && <Link to="/student/books">Explore Books</Link>}
      {isStudentLogged && <Link to="/student/home">Home</Link>}


      {(isStudentLogged || isLibrarianLogged) && (
        <button onClick={handleUserLogout} style={{ marginLeft: "10px", cursor: "pointer" }}>
          Logout
        </button>
      )}
    </nav >
  );
}
