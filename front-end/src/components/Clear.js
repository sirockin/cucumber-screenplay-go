import React, { useState } from 'react';

function Clear() {
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');

  const handleClear = async () => {
    setMessage('');
    setError('');

    try {
      const response = await fetch('/clear', {
        method: 'DELETE',
      });

      if (response.ok) {
        setMessage('All data cleared successfully!');
      } else {
        setError('Failed to clear data');
      }
    } catch (err) {
      setError(`Network error: ${err.message}`);
    }
  };

  return (
    <div>
      <h2>Admin: Clear All Data</h2>
      <p>This will clear all accounts and projects from the system.</p>

      {message && <div className="success">{message}</div>}
      {error && <div className="error">{error}</div>}

      <button onClick={handleClear} style={{ backgroundColor: '#dc3545', color: 'white', padding: '10px 20px', border: 'none', borderRadius: '3px', cursor: 'pointer' }}>
        Clear All Data
      </button>
    </div>
  );
}

export default Clear;