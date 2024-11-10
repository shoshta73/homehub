import { useRef, useState } from "react";
import { baseUrl } from "../../../helpers/api";
import { useNavigate } from "react-router-dom";

function Create() {
  const inputRef = useRef<HTMLInputElement>(null);
  const textareaRef = useRef<HTMLTextAreaElement>(null);
  const [error, setError] = useState<string | null>(null);

  const navigate = useNavigate();

  const submit = async () => {
    if (inputRef.current === null) {
      return;
    }

    if (textareaRef.current === null) {
      return;
    }

    if (inputRef.current.value.length === 0) {
      setError("Empty title is not allowed");
      return;
    }

    if (textareaRef.current.value.length === 0) {
      setError("Empty pastes are not allowed");
      return;
    }

    const res = await fetch(baseUrl() + "/pastebin/create", {
      method: "POST",
      credentials: "include",
      body: JSON.stringify({
        title: inputRef.current.value,
        content: textareaRef.current.value,
      }),
    });

    switch (res.status) {
      case 401:
        navigate("/login");
        break;
      case 409:
        setError(await res.text());
        break;
    }
  };

  const handleChange = () => {
    setError(null);
  };

  return (
    <>
      <input ref={inputRef} onChange={handleChange} />
      <textarea ref={textareaRef} onChange={handleChange} />
      {error !== null && <>{error}</>}
      <button onClick={submit}>Create</button>
    </>
  );
}

export default Create;
