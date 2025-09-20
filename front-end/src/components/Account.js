import React, { useState, useEffect } from 'react';
import { useParams, Link } from 'react-router-dom';

function Account() {
  const { name } = useParams();
  const [account, setAccount] = useState(null);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchAccount = async () => {
      try {
        const response = await fetch(`/accounts/${name}`);
        if (response.ok) {
          const accountData = await response.json();
          setAccount(accountData);
        } else {
          setError(`Account not found: ${name}`);
        }
      } catch (err) {
        setError(`Network error: ${err.message}`);
      }
    };

    if (name) {
      fetchAccount();
    }
  }, [name]);

  if (error) {
    return <div className="error">{error}</div>;
  }

  if (!account) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h2>Account: {account.name}</h2>

      <div className="account-info">
        <p>
          <strong>Status:</strong>{' '}
          {account.activated && <span className="status-activated">Activated</span>}
          {!account.activated && <span>Not Activated</span>}
        </p>
        <p>
          <strong>Authentication:</strong>{' '}
          {account.authenticated && <span className="status-authenticated">Authenticated</span>}
          {!account.authenticated && <span>Not Authenticated</span>}
        </p>
      </div>

      <div>
        {!account.activated && (
          <Link to={`/activate/${name}`}>
            <button className="activate">Activate Account</button>
          </Link>
        )}
        <Link to={`/account/${name}/projects`} style={{ marginLeft: '10px' }}>
          <button>View Projects</button>
        </Link>
      </div>
    </div>
  );
}

export default Account;