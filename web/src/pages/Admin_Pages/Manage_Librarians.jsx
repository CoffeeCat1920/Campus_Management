import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom"

export default function Manage_Librarians() {
  const navigate = useNavigate();

  const [librarians, setLibrarian] = useState([])
  const [name, setName] = useState([])
  const [password, setPassword] = useState([])

  const fetchLibrarians = async () => {
    const res = await fetch(`/all_librarians`);
    const data = await res.json();
    setLibrarian(data);
  };

  const handleAdd = async () => {
    await fetch(`/librarian`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, password }),
    });
    setName('');
    setPassword('');
    fetchLibrarians();
  };

  const deleteLibrarian = async (id) => {
    try {
      const res = await fetch(`/librarian/${id}`, {
        method: "DELETE",
      });

      if (res.ok) {
        fetchLibrarians()
      } else {
        const text = await res.text();
        console.log(`Failed to delete librarian : ${text}`);
      }
    } catch (error) {
      console.error("Error deleting librarian:", error);
      alert("An error occurred while deleting the librarian.");
    }
  };

  useEffect(() => {
    fetchLibrarians();
    console.log(librarians)
  }, [])

  return (
    <div style={{ padding: "20px", fontFamily: "Arial, sans-serif" }}>
      <h2>Librarian</h2>

      <div style={{ marginBottom: "20px", display: "flex", alignItems: "center", gap: "10px" }}>
        <input
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          style={{ padding: "6px", fontSize: "14px" }}
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          style={{ padding: "6px", fontSize: "14px" }}
        />
        <button onClick={handleAdd} style={{ padding: "6px 12px", cursor: "pointer" }}>
          Add
        </button>
      </div>

      <ul style={{ listStyle: "none", padding: 0 }}>
        {librarians.map((librarian) => (
          <li
            key={librarian.uuid}
            style={{
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
              padding: "8px 0",
              borderBottom: "1px solid #ddd",
            }}
          >
            <div style={{ flex: 1 }}>
              <strong>{librarian.name}</strong>
            </div>
            <div style={{ display: "flex", gap: "10px" }}>
              <button style={{ padding: "4px 8px", cursor: "pointer" }} onClick={
                () => navigate(`/edit/librarian/${librarian.uuid}`)
              } >Edit</button>
              <button style={{ padding: "4px 8px", cursor: "pointer" }} onClick={() => deleteLibrarian(librarian.uuid)}>Delete</button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );


}
