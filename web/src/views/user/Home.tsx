import { useEffect, useState } from "react";
import { Outlet, useLocation, useNavigate } from "react-router-dom";
import { baseUrl } from "../../helpers/api";

const Stats = () => {
  const [createdPastes, setCreatedPastes] = useState(0);
  const [createdError, setCreatedError] = useState<string | null>(null);
  const [createdPastesLoaded, setCreatedPastesLoaded] = useState(false);

  useEffect(() => {
    const getCreatedPastes = async () => {
      try {
        const res = await fetch(baseUrl() + "/pastebin/created/count", {
          method: "GET",
          credentials: "include",
        });
        const data = await res.text();
        setCreatedPastes(parseInt(data, 10));
      } catch (err) {
        setCreatedError(err as string);
      } finally {
        setCreatedPastesLoaded(true);
      }
    };

    if (!createdPastesLoaded) getCreatedPastes();
  }, [createdPastesLoaded]);

  return (
    <>
      <h1>Stats</h1>
      <h3>Pastebin</h3>
      <p>Created</p>
      {createdError !== null ? <>{createdError}</> : <>{createdPastes}</>}
    </>
  );
};

function Home() {
  const navigate = useNavigate();
  const location = useLocation();

  useEffect(() => {
    const validate = async () => {
      const res = await fetch(baseUrl() + "/auth/validate", {
        method: "POST",
        credentials: "include",
      });
      if (res.status !== 200) {
        navigate("/login");
      }
    };

    validate();
  }, []);

  return (
    <div>
      <div>
        <button onClick={() => navigate("/home/pastebin")}>Pastebin</button>
      </div>
      <div>{location.pathname === "/home" ? <Stats /> : <Outlet />}</div>
    </div>
  );
}

export default Home;
