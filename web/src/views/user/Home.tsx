import { useNavigate } from "react-router-dom";

function Home() {
  const navigate = useNavigate();

  return (
    <div>
      <div>
        <button onClick={() => navigate("/pastebin")}>Pastebin</button>
      </div>
    </div>
  );
}

export default Home;
