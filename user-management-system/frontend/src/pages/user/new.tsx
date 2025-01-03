import React, { useState } from 'react';
import { useRouter } from 'next/router';
import axios from 'axios';

const NewUser = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [error, setError] = useState('');
  const router = useRouter();

  // Kullanıcı ekleme işlemi
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!name || !email) {
      setError('Name and email are required.');
      return;
    }

    try {
      await axios.post('http://localhost:8080/users/', { name, email });
      router.push('/'); // Ana sayfaya yönlendirme
    } catch (err) {
      console.error('Failed to create user:', err);
      setError('Failed to create user. Please try again.');
    }
  };

  return (
    <div className="container">
      <h1>Create New User</h1>
      {error && <p>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="name">Name:</label>
          <input
            type="text"
            id="name"
            placeholder="Enter user name"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            placeholder="Enter user email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <button type="submit">Create</button>
        <button type="button" className="back" onClick={() => router.push('/')}>
          Back
        </button>
      </form>
    </div>
  );
};

export default NewUser;
