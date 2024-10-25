import { Button } from "@/components/ui/button";
import { Menubar, MenubarMenu, MenubarTrigger } from "@/components/ui/menubar";
import useAppState from "@/store/state";
import { MoonIcon, SunIcon } from "lucide-react";

function HomeView() {
  const state = useAppState();

  return (
    <>
      <Menubar>
        <div className="flex-grow h-auto" id="menubar-spacer" />
        <MenubarMenu>
          <MenubarTrigger onClick={state.toggleMode}>
            {state.mode === "light" ? <MoonIcon /> : <SunIcon />}
          </MenubarTrigger>
        </MenubarMenu>
      </Menubar>
      <h1 className="text-3xl font-bold underline">Hello from HomeHub!</h1>
      <Button>Click me</Button>
    </>
  );
}

export default HomeView;
