import { Link } from 'react-router-dom';
import { useAuth } from "../AuthContext";

export default function Navbar() {

  const { isAdminLogged, isStudentLogged, isLibrarianLogged } = useAuth()

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

      {!(isStudentLogged || isLibrarianLogged) && <Link to="/login">Login</Link>}

      {isLibrarianLogged && <Link to="/librarian/manage/students">Manage Students</Link>}
      {isLibrarianLogged && <Link to="/librarian/manage/books">Manage Books</Link>}
    </nav >
  );
}
