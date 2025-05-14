import { useEffect, useState } from "react";
import { useParams, useNavigate, useLocation } from "react-router-dom";

export default function Edit_Students() {
  const { id } = useParams();
  const location = useLocation();

  const [requests, setRequests] = useState([]);
  const [borrows, setBorrows] = useState([]); // New state
  const [borrowCount, setBorrowCount] = useState(0);
  const [activeInputId, setActiveInputId] = useState(null);
  const [daysInput, setDaysInput] = useState("");
  const [totalFine, setTotalFine] = useState(0); // State for total fine

  const studentName = location.state?.studentName;

  const fetchRequests = async () => {
    try {
      const res = await fetch(`/requests/${id}`);
      if (!res.ok) throw new Error("Failed to fetch students requests");
      const data = await res.json();
      setRequests(data);
    } catch (err) {
      console.error(err);
    }
  };

  const fetchBorrowCount = async () => {
    try {
      const res = await fetch(`/student/nob/${id}`);
      if (!res.ok) throw new Error("Failed to fetch borrow count");
      const data = await res.json();
      setBorrowCount(data);
    } catch (err) {
      console.error(err);
    }
  };

  const fetchBorrows = async () => {
    try {
      const res = await fetch(`/borrow/${id}`);
      if (!res.ok) throw new Error("Failed to fetch borrows");
      const data = await res.json();
      setBorrows(data);
    } catch (err) {
      console.error(err);
    }
  };

  const fetchTotalFine = async () => {
    try {
      const res = await fetch(`/librarian/fine/${id}`, {
        method: "PATCH",
      });
      if (!res.ok) throw new Error("Failed to fetch total fine");
      const data = await res.json();
      setTotalFine(data); // Assume the API returns { fine: totalFine }
      console.log(totalFine)
    } catch (err) {
      console.error(err);
    }
  };


  const handleReturn = async (borrowId) => {
    try {
      const res = await fetch(`/return_book/${borrowId}`, {
        method: "POST",
      });
      if (!res.ok) throw new Error("Failed to return book");
      fetchBorrows();  // Refresh the borrow list
      fetchBorrowCount(); // Refresh borrow count
    } catch (err) {
      console.error(err);
    }
  };

  const handleAccept = async (requestid, days) => {
    try {
      const res = await fetch(`/accept_request/${requestid}`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ days }),
      });
      if (!res.ok) throw new Error("Failed to accept request");

      setActiveInputId(null);
      setDaysInput("");
      fetchRequests();
      fetchBorrowCount();
      fetchBorrows(); // update borrows as well
    } catch (err) {
      console.error(err);
    }
  };

  const handleDecline = async (requestid) => {
    try {
      const res = await fetch(`/decline_request/${requestid}`, {
        method: "POST",
      });
      if (!res.ok) throw new Error("Failed to decline request");
      fetchRequests();
    } catch (err) {
      console.error(err);
    }
  };

  const handleSubmit = (requestId) => {
    const parsed = parseInt(daysInput, 10);
    if (isNaN(parsed)) {
      alert("Please enter a valid number");
      return;
    }
    handleAccept(requestId, parsed);
  };

  const handleCancel = () => {
    setActiveInputId(null);
    setDaysInput("");
  };

  const handleClearFine = async () => {
    try {
      const res = await fetch(`/librarian/clear_fine/${id}`, {
        method: "PATCH",
      });
      if (!res.ok) throw new Error("Failed to clear fine");
      setTotalFine(0); // Reset the fine in the UI
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    fetchRequests();
    fetchBorrowCount();
    fetchBorrows();
    fetchTotalFine(); // Fetch the total fine on component mount
  }, [id]);

  return (
    <div>
      <h2>
        Student: {studentName}
        <span style={{ marginLeft: "20px", fontSize: "0.9em", color: "gray" }}>
          Rented Books: {borrowCount}
        </span>
      </h2>
      {borrowCount >= 3 && (
        <p style={{ color: "red", fontWeight: "bold" }}>
          Max borrow limit reached (3)
        </p>
      )}

      {totalFine > 0 && (
        <div style={{ marginTop: "10px", color: "red", fontWeight: "bold" }}>
          <h4>Total Fine: Rs. {totalFine}</h4>
          <button onClick={handleClearFine}>Clear Fine</button>
        </div>
      )}

      <h4>Pending Requests</h4>
      <ul style={{ listStyle: "none", padding: 0 }}>
        {requests.map((request) => (
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

            {activeInputId === request.uuid ? (
              <div style={{ display: "flex", gap: "8px", alignItems: "center" }}>
                <input
                  type="number"
                  value={daysInput}
                  onChange={(e) => setDaysInput(e.target.value)}
                  onKeyDown={(e) => {
                    if (e.key === "Enter") handleSubmit(request.uuid);
                    if (e.key === "Escape") handleCancel();
                  }}
                  placeholder="Days"
                  style={{ width: "60px" }}
                />
                <button onClick={() => handleSubmit(request.uuid)}>OK</button>
                <button onClick={handleCancel}>Cancel</button>
              </div>
            ) : (
              <div style={{ display: "flex", gap: "10px" }}>
                {borrowCount < 3 && (
                  <button
                    style={{ padding: "4px 8px", cursor: "pointer" }}
                    onClick={() => setActiveInputId(request.uuid)}
                  >
                    Accept
                  </button>
                )}
                <button
                  style={{ padding: "4px 8px", cursor: "pointer" }}
                  onClick={() => handleDecline(request.uuid)}
                >
                  Deny
                </button>
              </div>
            )}
          </li>
        ))}
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
                display: "flex",
                justifyContent: "space-between",
                alignItems: "center"
              }}
            >
              <div>
                <div><strong>{borrow.book_name}</strong></div>
                {borrow.fine > 0 && <div style={{ color: "red" }}>Fine: {borrow.fine}</div>}
                <div><strong>Borrow Date:</strong> {borrow.borrow_date}</div>
                <div><strong>Return By:</strong> {borrow.return_date}</div>
              </div>

              <button
                style={{ padding: "6px 12px", cursor: "pointer" }}
                onClick={() => handleReturn(borrow.uuid)}
              >
                Return
              </button>
            </li>
          ))
        )}
      </ul>

    </div>
  );
}
