import React, { useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';

function Activate() {
  const { name } = useParams();
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const handleActivate = async () => {
    setMessage('');
    setError('');

    try {
      const response = await fetch(`/accounts/${name}/activate`, {
        method: 'POST',
      });

      if (response.ok) {
        setMessage(`Account ${name} activated successfully!`);
        setTimeout(() => {
          navigate(`/account/${name}`);
        }, 1500);
      } else {
        const errorData = await response.text();
        setError(`Failed to activate account: ${errorData}`);
      }
    } catch (err) {
      setError(`Network error: ${err.message}`);
    }
  };

  return (
    <div>
      <h2>Activate Account</h2>
      <p>Activate account for: <strong>{name}</strong></p>

      {message && <div className="success">{message}</div>}
      {error && <div className="error">{error}</div>}

      <button className="activate" onClick={handleActivate}>
        Activate Account
      </button>
    </div>
  );
}

export default Activate;