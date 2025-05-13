import { useEffect, useState } from "react";

export default function Explore_Books() {

  const [books, setBooks] = useState([]);

  const fetchBooks = async () => {
    const res = await fetch(`/all_books`);
    const data = await res.json();
    setBooks(data);
  };

  useEffect(() => {
    fetchBooks();
  }, []);

  const handleRequest = async (isbn) => {
    try {
      const res = await fetch(`/request`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ isbn }),
      });
      if (!res.ok) {
        throw new Error("Failed to make a request");
      }
    } catch (err) {
      console.error(err);
    }
  }

  return (
    <div style={{ padding: "20px", fontFamily: "Arial, sans-serif" }}>
      <h2>Books</h2>

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
              <button style={{ padding: "4px 8px", cursor: "pointer" }} onClick={() => handleRequest(book.isbn)}>Request</button>
            </div>
          </li>
        ))}
      </ul>
    </div>

  );
}
