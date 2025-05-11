import { useEffect, useState } from "react"

export default function Admin_Dashborad() {

  const [isLoggedIn, setIsLoggedIn] = useState(null)

  useEffect(() => {
    fetch('/admin/data', {
      method: 'GET',
      credentials: 'include'
    }).then((res) => {
      if (res.status == 200) {
        setIsLoggedIn(true);
      } else if (res.status == 401) {
        setIsLoggedIn(false);
      } else {
        throw new Error(`Unexpected status: ${res.status}`);
      }
    }).catch((err) => {
      console.error('Error checking login status:', err);
      setIsLoggedIn(false);
    })
  }, []);

  return (
    <div>
      {isLoggedIn === null && <p>Checking login status...</p>}
      {isLoggedIn === true && <p>✅ Admin is logged in</p>}
      {isLoggedIn === false && <p>❌ Admin not logged in</p>}
    </div>
  )
}
