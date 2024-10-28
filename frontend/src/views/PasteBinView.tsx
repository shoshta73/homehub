import Menubar from "@/components/menubar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { useRef } from "react";

function PasteBinView() {
  const titleRef = useRef<HTMLInputElement>(null);
  const textAreaRef = useRef<HTMLTextAreaElement>(null);

  const handleSubmit = () => {
    if (textAreaRef.current == null) {
      return;
    }

    if (titleRef.current == null) {
      return;
    }
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
