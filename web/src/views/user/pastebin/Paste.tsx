import { useEffect, useState } from "react";
import { baseUrl } from "../../../helpers/api";
import { useParams } from "react-router-dom";

type Paste = {
  id: string;
  title: string;
  created: Date;
  updated: Date;
  content: string;
};

function Paste() {
  const { id } = useParams();
  const [paste, setPaste] = useState<Paste>();
  const [error, setError] = useState<string | null>(null);
  const [loaded, setLoaded] = useState(false);

  useEffect(() => {
    const getPaste = async () => {
      try {
        const res = await fetch(baseUrl() + `/pastebin/paste/${id}`, { method: "GET", credentials: "include" });

        const data = await res.json();
        setPaste(data);
      } catch (err) {
        setError(err as string);
      } finally {
        setLoaded(true);
      }
    };

    if (!loaded) getPaste();
  });
  return (
    <>
      {error !== null && <>{error}</>}
      {paste !== undefined && (
        <>
          <h1>{paste.title}</h1>
          <p>{paste.content}</p>
        </>
      )}
    </>
  );
}

export default Paste;
