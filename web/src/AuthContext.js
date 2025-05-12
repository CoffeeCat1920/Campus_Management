import { createContext, useContext, useState, useEffect } from "react";

const AuthContext = createContext();

export function AuthProvider({ children }) {
  const [isAdminLogged, setAdminLogged] = useState(false);
  const [isStudentLogged, setStudentLogged] = useState(false);
  const [isLibrarianLogged, setLibrarianLogged] = useState(false);

  useEffect(() => {
    fetch('/admin/data', {
      method: 'GET',
      credentials: 'include'
    })
      .then(res => {
        if (res.status === 200) setAdminLogged(true);
        else setAdminLogged(false);
      })
      .catch(() => setAdminLogged(false));


    fetch('/librarian/data', {
      method: 'GET',
      credentials: 'include'
    })
      .then(res => {
        if (res.status === 200) setLibrarianLogged(true);
        else setLibrarianLogged(false);
      })
      .catch(() => setLibrarianLogged(false));

    fetch('/student/data', {
      method: 'GET',
      credentials: 'include'
    })
      .then(res => {
        if (res.status === 200) setStudentLogged(true);
        else setStudentLogged(false);
      })
      .catch(() => setStudentLogged(false));
  }, []);

  return (
    <AuthContext.Provider value={{ isAdminLogged, isStudentLogged, isLibrarianLogged }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  return useContext(AuthContext);
}
