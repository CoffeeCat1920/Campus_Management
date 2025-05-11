import { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

export default function Navbar() {

  const [isAdminLogged, setAdminLogged] = useState(false)
  const [isUserLogged, setUserLogged] = useState(false)

  useEffect(() => {
    fetch('/admin/data', {
      method: 'GET',
      credentials: 'include'
    }).then((res) => {
      if (res.status == 200) {
        setAdminLogged(true);
      } else if (res.status == 401) {
        setAdminLogged(false);
      } else {
        throw new Error(`Unexpected status: ${res.status}`);
      }
    }).catch((err) => {
      console.error('Error checking login status:', err);
      setAdminLogged(false);
    })
  }, []);

  return (
    <nav>
      {isAdminLogged ? (
        <Link to="/admin/dashboard">Home</Link>
      ) : (
        <Link to="/">Home</Link>
      )}
      {isAdminLogged && <Link to="/manage/students">Manage Students</Link>}
      {!isAdminLogged && <Link to="/login">Login</Link>}
      {!isAdminLogged && <Link to="/admin/login">Admin Login</Link>}
    </nav >
  );
}
