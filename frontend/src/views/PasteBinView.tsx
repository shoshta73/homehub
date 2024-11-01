import Menubar from "@/components/menubar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";

function PasteBinView() {
  const titleRef = useRef<HTMLInputElement>(null);
  const textAreaRef = useRef<HTMLTextAreaElement>(null);

  const navigate = useNavigate();

  useEffect(() => {
    const validate = async () => {
      try {
        const res = await fetch("/auth/validate", {
          method: "POST",
          credentials: "include",
        });
        if (res.status !== 200) {
          navigate("/login");
        }
      } catch (err) {
        console.error(err);
        navigate("/login");
      }
    };

    validate();
  }, []);

  const handleSubmit = () => {
    if (textAreaRef.current == null) {
      return;
    }

    if (titleRef.current == null) {
      return;
    }

    fetch(`${import.meta.env.VITE_API_URL ?? ""}/pastebin/create`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        title: titleRef.current.value,
        content: textAreaRef.current.value,
      }),
    });
  };

  return (
    <>
      <Menubar />
      <div className="flex flex-col flex-grow w-2/3 m-auto mt-4 gap-2">
        <Input ref={titleRef} placeholder="Pastebin name" />
        <Textarea ref={textAreaRef} placeholder="Paste your text here..." />
        <Button onClick={handleSubmit}>Create Paste</Button>
      </div>
    </>
  );
}

export default PasteBinView;
