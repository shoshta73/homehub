import { useNavigate } from "react-router-dom";

function App() {
  const navigate = useNavigate();

  return (
    <>
      <h1>Hello From Homehub</h1>
      <button onClick={() => navigate("/")}>Go to home</button>
      <button onClick={() => navigate("/register")}>Register</button>
      <button onClick={() => navigate("/login")}>Login</button>
    </>
  );
}

export default App;
