import { useRef, useState } from "react";
import { baseUrl } from "../helpers/api";
import { useNavigate } from "react-router-dom";

function Register() {
  const [error, setError] = useState<string | null>(null);

  const nameRef = useRef<HTMLInputElement | null>(null);
  const usernameRef = useRef<HTMLInputElement | null>(null);
  const emailRef = useRef<HTMLInputElement | null>(null);
  const passwordRef = useRef<HTMLInputElement | null>(null);
  const confirmPasswordRef = useRef<HTMLInputElement | null>(null);

  const navigate = useNavigate();

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (
      usernameRef.current === null ||
      emailRef.current === null ||
      passwordRef.current === null ||
      confirmPasswordRef.current === null
    ) {
      return;
    }

    if (usernameRef.current.value.length < 3) {
      setError("Username is too short");
      return;
    }

    if (passwordRef.current.value.length < 8) {
      setError("Password is too short");
      return;
    }

    if (passwordRef.current.value !== confirmPasswordRef.current.value) {
      setError("Passwords do not match");
      return;
    }

    fetch(baseUrl() + "/auth/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body:
        nameRef.current !== null
          ? JSON.stringify({
              name: nameRef.current.value,
              username: usernameRef.current.value,
              email: emailRef.current.value,
              password: passwordRef.current.value,
            })
          : JSON.stringify({
              name: "",
              username: usernameRef.current.value,
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
      <h1>Register</h1>
      <label htmlFor="name">Name</label>
      <input ref={nameRef} type="text" id="name" name="name" placeholder="Name" autoComplete="name" />
      <label htmlFor="username">Username</label>
      <input
        ref={usernameRef}
        type="text"
        id="username"
        name="username"
        placeholder="Username"
        autoComplete="username"
      />
      <label htmlFor="email">Email</label>
      <input ref={emailRef} type="email" id="email" name="email" placeholder="Email" autoComplete="email" />
      <label htmlFor="password">Password</label>
      <input
        ref={passwordRef}
        type="password"
        id="password"
        name="password"
        placeholder="Password"
        autoComplete="new-password"
      />
      <label htmlFor="confirm-password">Confirm Password</label>
      <input
        ref={confirmPasswordRef}
        id="confirm-password"
        type="password"
        name="confirm-password"
        placeholder="Confirm Password"
        autoComplete="new-password"
      />
      {error !== null && <div>{error}</div>}
      <button type="submit">Register</button>
    </form>
  );
}

export default Register;
