import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { baseUrl } from "../../helpers/api";

function Home() {
  const navigate = useNavigate();

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
        <button onClick={() => navigate("/pastebin")}>Pastebin</button>
      </div>
    </div>
  );
}

export default Home;
