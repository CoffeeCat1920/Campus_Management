import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Manage_Books() {
  const navigate = useNavigate();

  const [books, setBooks] = useState([]);
  const [name, setName] = useState('');
  const [isbn, setISBN] = useState('');

  const fetchBooks = async () => {
    const res = await fetch(`/all_books`);
    const data = await res.json();
    setBooks(data);
  };

  const handleAdd = async () => {
    await fetch(`/books`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ isbn, name }),
    });
    setName('');
    setISBN('');
    fetchBooks();
  };

  const handleAvailableToggle = async (id) => {
    try {
      const res = await fetch(`/toggle_books/${id}`, {
        method: "PATCH",
      });

      if (res.ok) {
        fetchBooks();
      } else {
        const text = await res.text();
        console.log(`Failed to toggle book: ${text}`);
      }
    } catch (error) {
      console.error("Error deleting book:", error);
      alert("An error occurred while toggling the book.");
    }

  }

  const deleteBook = async (id) => {
    try {
      const res = await fetch(`/books/${id}`, {
        method: "DELETE",
      });

      if (res.ok) {
        fetchBooks();
      } else {
        const text = await res.text();
        console.log(`Failed to delete book: ${text}`);
      }
    } catch (error) {
      console.error("Error deleting book:", error);
      alert("An error occurred while deleting the book.");
    }
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  return (
    <div style={{ padding: "20px", fontFamily: "Arial, sans-serif" }}>
      <h2>Books</h2>

      <div style={{ marginBottom: "20px", display: "flex", alignItems: "center", gap: "10px" }}>
        <input
          placeholder="Book Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          style={{ padding: "6px", fontSize: "14px" }}
        />
        <input
          placeholder="ISBN"
          value={isbn}
          onChange={(e) => setISBN(e.target.value)}
          style={{ padding: "6px", fontSize: "14px" }}
        />
        <button onClick={handleAdd} style={{ padding: "6px 12px", cursor: "pointer" }}>
          Add
        </button>
      </div>

      <ul style={{ listStyle: "none", padding: 0 }}>
        {books.map((book) => (
          <li
            key={book.uuid}
            style={{
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
              padding: "8px 0",
              borderBottom: "1px solid #ddd",
            }}
          >
            <div style={{ flex: 1 }}>
              <div>
                <p><strong>📕 {book.name}</strong></p>
                <p>ISBN: {book.isbn}</p>
                <p>Available: {book.available ? 'Yes' : 'No'}</p>
              </div>
            </div>
            <div style={{ display: "flex", gap: "10px" }}>
              <button
                style={{ padding: "4px 8px", cursor: "pointer" }}
                onClick={() => navigate(`/librarian/edit/book/${book.uuid}`)}
              >
                Edit
              </button>
              <button
                style={{ padding: "4px 8px", cursor: "pointer" }}
                onClick={() => deleteBook(book.uuid)}
              >
                Delete
              </button>
              <button
                style={{ padding: "4px 8px", cursor: "pointer" }}
                onClick={() => handleAvailableToggle(book.uuid)}
              >
                Toggle Available
              </button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}
