import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";

export default function Edit_Librarian() {
  const { id } = useParams();
  const [name, setName] = useState('')
  const [newPassword, setNewPassword] = useState('')

  const navigate = useNavigate();

  const handleEdit = async () => {
    await fetch(`/librarian/${id}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, new_password: newPassword }),
    });
    navigate(`/manage/librarians`)
  };

  const fetchLibrarian = async () => {
    try {
      const res = await fetch(`/librarian/${id}`);
      if (!res.ok) {
        throw new Error("Failed to fetch librarian");
      }
      const data = await res.json();
      setName(data.name)
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    fetchLibrarian();
  }, [id]);

  return (
    <div>

      <h2>Edit the user</h2>

      <div style={{ marginBottom: "20px", display: "flex", alignItems: "center", gap: "10px" }}>
        <input
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          style={{ padding: "6px", fontSize: "14px" }}
        />
        <input
          type="password"
          placeholder="New Password"
          value={newPassword}
          onChange={(e) => setNewPassword(e.target.value)}
          style={{ padding: "6px", fontSize: "14px" }}
        />
        <button style={{ padding: "4px 8px", cursor: "pointer" }} onClick={() => handleEdit()}>Change</button>
      </div>
    </div>
  );
}
