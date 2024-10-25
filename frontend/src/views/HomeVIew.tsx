import { Button } from "@/components/ui/button";
import { Menubar, MenubarMenu } from "@/components/ui/menubar";
import useAppState from "@/store/state";
import { LogIn, MoonIcon, SunIcon, UserIcon } from "lucide-react";
import { useNavigate } from "react-router-dom";

function HomeView() {
  const state = useAppState();
  const navigate = useNavigate();

  return (
    <>
      <Menubar className="h-14">
        <div className="flex-grow h-auto" id="menubar-spacer" />

        <MenubarMenu>
          <>
            {/* Register and Login buttons */}
            <Button
              className="my-1 px-1 py-1 h-12 flex flex-row text-base [&_svg]:size-6"
              variant={"outline"}
              onClick={() => navigate("/register")}
            >
              <UserIcon />
              <p>Register</p>
            </Button>
            <Button
              className="my-1 px-1 py-1 h-12 flex flex-row text-base [&_svg]:size-6"
              variant={"outline"}
              onClick={() => navigate("/login")}
            >
              <LogIn />
              <p>Login</p>
            </Button>
          </>
        </MenubarMenu>

        <MenubarMenu>
          <Button
            className="my-1 px-1 py-1 w-12 h-12 flex flex-row text-base [&_svg]:size-6"
            variant={"outline"}
            onClick={state.toggleMode}
          >
            {state.mode === "light" ? <MoonIcon /> : <SunIcon />}
          </Button>
        </MenubarMenu>
      </Menubar>
      <h1 className="text-3xl font-bold underline">Hello from HomeHub!</h1>
      <Button>Click me</Button>
    </>
  );
}

export default HomeView;
