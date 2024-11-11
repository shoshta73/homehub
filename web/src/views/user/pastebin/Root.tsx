import { useEffect, useState } from "react";
import { baseUrl } from "../../../helpers/api";
import { useNavigate } from "react-router-dom";

type Paste = { id: string; title: string; created: Date; updated: Date };

function Root() {
  const [createdPastes, setCreatedPastes] = useState<Paste[]>([]);
  const [createdPastesError, setCreatedPastesError] = useState<string | null>(null);
  const [createdPastesLoaded, setCreatedPastesLoaded] = useState(false);

  const navigate = useNavigate();

  useEffect(() => {
    const getCreatePastes = async () => {
      try {
        const res = await fetch(baseUrl() + "/pastebin/created/all", { method: "GET", credentials: "include" });

        const data = (await res.json()) as {
          pastes: Paste[];
        };
        setCreatedPastes(data.pastes);
      } catch (err) {
        setCreatedPastesError(err as string);
      } finally {
        setCreatedPastesLoaded(true);
      }
    };

    if (!createdPastesLoaded) getCreatePastes();
  }, []);

  return (
    <>
      <h1>Created</h1>
      {!createdPastesLoaded && <>Pastes are loading</>}
      {createdPastesLoaded && createdPastesError !== null ? (
        <>{createdPastesError}</>
      ) : (
        <>
          {createdPastes.map((paste) => {
            return (
              <div key={paste.id} onClick={() => navigate(`/home/pastebin/paste/${paste.id}`)}>
                <h1>{paste.title}</h1>
              </div>
            );
          })}
        </>
      )}
    </>
  );
}

export default Root;
