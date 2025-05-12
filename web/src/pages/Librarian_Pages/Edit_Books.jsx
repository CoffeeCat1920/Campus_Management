import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";

export default function Edit_Student() {
  const { id } = useParams();
  const navigate = useNavigate();

  const [name, setName] = useState('');
  const [ISBN, setISBN] = useState('');

  const handleEdit = async () => {
    await fetch(`/books/${id}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, ISBN }),
    });
    navigate('librarian/manage/books');
  };

  const fetchBook = async () => {
    try {
      const res = await fetch(`/books/${id}`);
      if (!res.ok) throw new Error("Failed to fetch books");

      const data = await res.json();
      setName(data.name);
      setISBN(data.isbn)
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    fetchBook();
  }, [id]);

  return (
    <div style={{ padding: "20px", fontFamily: "Arial, sans-serif" }}>
      <h2>Edit Books</h2>
      <div style={{ marginBottom: "20px", display: "flex", alignItems: "center", gap: "10px" }}>
        <input
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          style={{ padding: "6px", fontSize: "14px" }}
        />
        <input
          placeholder="ISBN"
          value={ISBN}
          onChange={(e) => setISBN(e.target.value)}
          style={{ padding: "6px", fontSize: "14px" }}
        />
        <button onClick={handleEdit} style={{ padding: "6px 12px", cursor: "pointer" }}>
          Save
        </button>
      </div>
    </div>
  );
}
