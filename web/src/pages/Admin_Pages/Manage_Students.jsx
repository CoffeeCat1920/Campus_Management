import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom"

export default function Manage_Students() {

  const navigate = useNavigate();

  const [students, setStudents] = useState([])
  const [name, setName] = useState([])
  const [password, setPassword] = useState([])

  const fetchStudents = async () => {
    const res = await fetch(`/all_students`);
    const data = await res.json();
    setStudents(data);
  };

  const handleAdd = async () => {
    await fetch(`/student`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, password }),
    });
    setName('');
    setPassword('');
    fetchStudents();
  };

  const deleteStudent = async (id) => {
    try {
      const res = await fetch(`/student/${id}`, {
        method: "DELETE",
      });

      if (res.ok) {
        fetchStudents();
      } else {
        const text = await res.text();
        console.log(`Failed to delete student: ${text}`);
      }
    } catch (error) {
      console.error("Error deleting student:", error);
      alert("An error occurred while deleting the student.");
    }
  };

  useEffect(() => {
    fetchStudents();
    console.log(students)
  }, [])

  return (
    <div style={{ padding: "20px", fontFamily: "Arial, sans-serif" }}>
      <h2>Students</h2>

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
        {students.map((student) => (
          <li
            key={student.uuid}
            style={{
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
              padding: "8px 0",
              borderBottom: "1px solid #ddd",
            }}
          >
            <div style={{ flex: 1 }}>
              <strong>👤{student.name}</strong> — Rented Books: {student.rentedbooks}
            </div>
            <div style={{ display: "flex", gap: "10px" }}>
              <button style={{ padding: "4px 8px", cursor: "pointer" }} onClick={
                () => navigate(`/edit/student/${student.uuid}`)
              } >Edit</button>
              <button style={{ padding: "4px 8px", cursor: "pointer" }} onClick={() => deleteStudent(student.uuid)}>Delete</button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );

}
