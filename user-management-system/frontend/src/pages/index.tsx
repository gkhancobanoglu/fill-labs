import React, { useEffect, useState } from 'react';
import { useRouter } from 'next/router';
import axios from 'axios';

// User interface
interface User {
  id: number;
  name: string;
  email: string;
}

const Home = () => {
  const [users, setUsers] = useState<User[]>([]);
  const router = useRouter();

  // Fetch Users
  useEffect(() => {
    const fetchUsers = async () => {
      try {
        // Specify the type of response data as User[]
        const response = await axios.get<User[]>('http://localhost:8080/users/');
        setUsers(response.data); // Now TypeScript knows the type of response.data
      } catch (error) {
        console.error('Failed to fetch users:', error);
      }
    };
    fetchUsers();
  }, []);

  // User Delete
  const handleDelete = async (id: number) => {
    try {
      await axios.delete(`http://localhost:8080/users/${id}`);
      setUsers(users.filter((user) => user.id !== id));
    } catch (error) {
      console.error('Failed to delete user:', error);
    }
  };

  return (
    <div className="container">
      <h1>User Management</h1>
      <button onClick={() => router.push('/user/new')}>New User</button>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Email</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          {users.map((user) => (
            <tr key={user.id}>
              <td>{user.id}</td>
              <td>{user.name}</td>
              <td>{user.email}</td>
              <td>
                <button onClick={() => router.push(`/user/${user.id}`)}>Edit</button>
                <button className="delete" onClick={() => handleDelete(user.id)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default Home;
