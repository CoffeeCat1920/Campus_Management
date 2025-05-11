import React, { useState, useEffect } from "react";

function BookManager() {
  const [books, setBooks] = useState([]);
  const [isbn, setIsbn] = useState("");
  const [name, setName] = useState("");
  const [editName, setEditName] = useState({});

  const fetchBooks = async () => {
    const res = await fetch(`/all_books`);
    const data = await res.json();
    setBooks(data);
  };

  const handleAdd = async () => {
    if (!isbn || !name) return;
    await fetch(`/books`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ isbn, name, available: true }),
    });
    setIsbn("");
    setName("");
    fetchBooks();
  };

  const handleDelete = async (isbn) => {
    await fetch(`/books/${isbn}`, { method: "DELETE" });
    fetchBooks();
  };

  const handleToggle = async (isbn) => {
    await fetch(`/toggle_books/${isbn}`, { method: "PATCH" });
    fetchBooks();
  };

  const handleEdit = async (isbn) => {
    if (!editName[isbn]) return;
    await fetch(`/books/${isbn}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name: editName[isbn] }),
    });
    setEditName({ ...editName, [isbn]: "" });
    fetchBooks();
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h2>Book Manager</h2>

      <div style={{ marginBottom: "10px" }}>
        <input
          placeholder="ISBN"
          value={isbn}
          onChange={(e) => setIsbn(e.target.value)}
        />
        <input
          placeholder="Book Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          style={{ marginLeft: "5px" }}
        />
        <button onClick={handleAdd} style={{ marginLeft: "5px" }}>
          Add
        </button>
      </div>

      {books.length === 0 ? (
        <p>No books found.</p>
      ) : (
        <ul>
          {books.map((book) => (
            <li key={book.ISBN} style={{ marginBottom: "8px" }}>
              <span>
                <b>{book.Name}</b> ({book.ISBN}) - Available:{" "}
                {book.Available ? "Yes" : "No"}
              </span>
              <br />
              <input
                placeholder="Edit Name"
                value={editName[book.ISBN] || ""}
                onChange={(e) =>
                  setEditName({ ...editName, [book.ISBN]: e.target.value })
                }
              />
              <button onClick={() => handleEdit(book.ISBN)}>Edit</button>
              <button onClick={() => handleToggle(book.ISBN)}>Toggle</button>
              <button onClick={() => handleDelete(book.ISBN)}>Delete</button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default BookManager;
