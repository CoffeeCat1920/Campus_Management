import React, { useState } from 'react';

function LoginForm() {
  const [formData, setFormData] = useState({ name: '', password: '' });
  const [message, setMessage] = useState('');

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setMessage('');

    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        credentials: 'include', // important for setting cookies
        body: JSON.stringify(formData)
      });

      if (response.ok) {
        setMessage('Login successful!');
        // Redirect or update app state here if needed
      } else {
        const errorText = await response.text();
        setMessage(`Login failed: ${errorText}`);
      }
    } catch (error) {
      console.error('Error:', error);
      setMessage('An error occurred. Try again.');
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Login</h2>

      <label>
        Name:
        <br />
        <input
          type="text"
          name="name"
          value={formData.name}
          onChange={handleChange}
          required
        />
      </label>
      <br />

      <label>
        Password:
        <br />
        <input
          type="password"
          name="password"
          value={formData.password}
          onChange={handleChange}
          required
        />
      </label>
      <br />

      <button type="submit">Login</button>
      <p>{message}</p>
    </form >
  );
}

export default LoginForm;
