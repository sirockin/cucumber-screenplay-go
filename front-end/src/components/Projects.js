import React, { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';

function Projects() {
  const { name } = useParams();
  const [projects, setProjects] = useState([]);
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');

  const fetchProjects = async () => {
    try {
      const response = await fetch(`/accounts/${name}/projects`);
      if (response.ok) {
        const projectsData = await response.json();
        setProjects(projectsData || []);
      } else {
        setError(`Failed to load projects for ${name}`);
      }
    } catch (err) {
      setError(`Network error: ${err.message}`);
    }
  };

  useEffect(() => {
    if (name) {
      fetchProjects();
    }
  }, [name]);

  const handleCreateProject = async () => {
    setMessage('');
    setError('');

    try {
      const response = await fetch(`/accounts/${name}/projects`, {
        method: 'POST',
      });

      if (response.ok) {
        setMessage('Project created successfully!');
        // Refresh the projects list
        setTimeout(() => {
          fetchProjects();
          setMessage('');
        }, 1500);
      } else {
        const errorData = await response.text();
        setError(`Failed to create project: ${errorData}`);
      }
    } catch (err) {
      setError(`Network error: ${err.message}`);
    }
  };

  return (
    <div>
      <h2>Projects for {name}</h2>

      {message && <div className="project-created">{message}</div>}
      {error && <div className="error">{error}</div>}

      <button className="create-project" onClick={handleCreateProject}>
        Create New Project
      </button>

      <div className="projects-list">
        {projects.length === 0 ? (
          <p>No projects found.</p>
        ) : (
          <ul>
            {projects.map((project, index) => (
              <li key={index} className="project-item">
                Project {index + 1}
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
}

export default Projects;