import { useEffect, useState } from "react";

export default function Edit_Students() {
  const [studentId, setStudentId] = useState(null);
  const [requests, setRequests] = useState([]);
  const [borrows, setBorrows] = useState([]);
  const [borrowCount, setBorrowCount] = useState(0);
  const [totalFine, setTotalFine] = useState(0);
  const [loading, setLoading] = useState(true);

  const fetchStudentData = async () => {
    try {
      const res = await fetch("/student/data", {
        method: "GET",
        credentials: "include"
      });

      if (!res.ok) throw new Error("Failed to fetch student data");

      const data = await res.json();
      setStudentId(data.UUID);
      return data.UUID;
    } catch (err) {
      console.error("Error fetching student data:", err);
      return null;
    }
  };

  const fetchRequests = async (id) => {
    if (!id) return;

    try {
      const res = await fetch(`/requests/${id}`);
      if (!res.ok) throw new Error("Failed to fetch student requests");
      const data = await res.json();
      setRequests(data);
    } catch (err) {
      console.error("Error fetching requests:", err);
    }
  };

  const fetchBorrowCount = async (id) => {
    if (!id) return;

    try {
      const res = await fetch(`/student/nob/${id}`);
      if (!res.ok) throw new Error("Failed to fetch borrow count");
      const data = await res.json();
      setBorrowCount(data);
    } catch (err) {
      console.error("Error fetching borrow count:", err);
    }
  };

  const fetchBorrows = async (id) => {
    if (!id) return;

    try {
      const res = await fetch(`/student/borrow/${id}`);
      if (!res.ok) throw new Error("Failed to fetch borrows");
      const data = await res.json();
      setBorrows(data);
    } catch (err) {
      console.error("Error fetching borrows:", err);
    }
  };

  const fetchTotalFine = async (id) => {
    if (!id) return;

    try {
      const res = await fetch(`/librarian/fine/${id}`, {
        method: "PATCH",
      });
      if (!res.ok) throw new Error("Failed to fetch total fine");
      const data = await res.json();
      setTotalFine(data);
    } catch (err) {
      console.error("Error fetching total fine:", err);
    }
  };

  // Initialize all data fetching
  useEffect(() => {
    const initializeData = async () => {
      setLoading(true);
      const uuid = await fetchStudentData();

      if (uuid) {
        await Promise.all([
          fetchRequests(uuid),
          fetchBorrowCount(uuid),
          fetchBorrows(uuid),
          fetchTotalFine(uuid)
        ]);
      }

      setLoading(false);
    };

    initializeData();
  }, []);

  if (loading) {
    return <div>Loading student data...</div>;
  }

  if (!studentId) {
    return <div>Error: Unable to retrieve student information. Please log in again.</div>;
  }

  return (
    <div>
      {borrowCount >= 3 && (
        <p style={{ color: "red", fontWeight: "bold" }}>
          Max borrow limit reached (3)
        </p>
      )}

      {totalFine > 0 && (
        <div style={{ marginTop: "10px", color: "red", fontWeight: "bold" }}>
          <h4>Total Fine: Rs. {totalFine}</h4>
        </div>
      )}

      <h4>Pending Requests</h4>
      <ul style={{ listStyle: "none", padding: 0 }}>
        {requests.length === 0 ? (
          <li>No pending requests.</li>
        ) : (
          requests.map((request) => (
            <li
              key={request.uuid}
              style={{
                display: "flex",
                alignItems: "center",
                justifyContent: "space-between",
                padding: "8px 0",
                borderBottom: "1px solid #ddd",
              }}
            >
              <div style={{ flex: 1 }}>
                <strong>{request.bookname}</strong>
              </div>
            </li>
          ))
        )}
      </ul>

      <h4 style={{ marginTop: "30px" }}>Borrowed Books</h4>
      <ul style={{ listStyle: "none", padding: 0 }}>
        {borrows.length === 0 ? (
          <li>No books currently borrowed.</li>
        ) : (
          borrows.map((borrow) => (
            <li
              key={borrow.uuid}
              style={{
                padding: "8px 0",
                borderBottom: "1px solid #ddd",
              }}
            >
              <div><strong>{borrow.book_name}</strong></div>
              {borrow.fine > 0 && <div style={{ color: "red" }}>Fine: {borrow.fine}</div>}
              <div><strong>Borrow Date:</strong> {borrow.borrow_date}</div>
              <div><strong>Return By:</strong> {borrow.return_date}</div>
            </li>
          ))
        )}
      </ul>
    </div>
  );
}
