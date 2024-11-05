import { useRef, useState } from "react";
import { baseUrl } from "../helpers/api";
import { useNavigate } from "react-router-dom";

function Login() {
  const [error, setError] = useState<string | null>(null);

  const emailRef = useRef<HTMLInputElement | null>(null);
  const passwordRef = useRef<HTMLInputElement | null>(null);

  const navigate = useNavigate();

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (emailRef.current === null || passwordRef.current === null) {
      return;
    }

    fetch(baseUrl() + "/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({
        email: emailRef.current.value,
        password: passwordRef.current.value,
      }),
    })
      .then(() => navigate("/home"))
      .catch(console.error);
  };

  return (
    <form
      onSubmit={handleSubmit}
      onChange={() => {
        setError(null);
      }}
    >
      <h1>Login</h1>
      <input ref={emailRef} type="email" id="email" name="email" placeholder="Email" autoComplete="email" />
      <label htmlFor="password">Password</label>
      <input
        ref={passwordRef}
        type="password"
        id="password"
        name="password"
        placeholder="Password"
        autoComplete="current-password"
      />
      {error !== null && <div>{error}</div>}
      <button type="submit">Login</button>
    </form>
  );
}

export default Login;
