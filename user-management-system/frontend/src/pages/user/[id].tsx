import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import axios from 'axios';


interface User {
  name: string;
  email: string;
}

const EditUser = () => {
  const router = useRouter();
  const { id } = router.query;
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [error, setError] = useState('');

  
  useEffect(() => {
    if (!id) return;

    const fetchUser = async () => {
      try {
        
        const response = await axios.get<User>(`http://localhost:8080/users/${id}`);
        setName(response.data.name);
        setEmail(response.data.email);
      } catch (err) {
        console.error('Failed to fetch user:', err);
        setError('Failed to fetch user. Please try again.');
      }
    };

    fetchUser();
  }, [id]);

  // User Update
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();

    if (!name || !email) {
      setError('Name and email are required.');
      return;
    }

    try {
      await axios.put(`http://localhost:8080/users/${id}`, { name, email });
      router.push('/'); 
    } catch (err) {
      console.error('Failed to update user:', err);
      setError('Failed to update user. Please try again.');
    }
  };

  return (
    <div className="container">
      <h1>Edit User</h1>
      {error && <p>{error}</p>}
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="name">Name:</label>
          <input
            type="text"
            id="name"
            placeholder="Enter new name"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>
        <div>
          <label htmlFor="email">Email:</label>
          <input
            type="email"
            id="email"
            placeholder="Enter new email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </div>
        <button type="submit">Save</button>
        <button type="button" className="back" onClick={() => router.push('/')}>
          Back
        </button>
      </form>
    </div>
  );
};

export default EditUser;
