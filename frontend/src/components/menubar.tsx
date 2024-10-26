import { Button } from "@/components/ui/button";
import { Menubar as MenubarPrimitive, MenubarMenu } from "@/components/ui/menubar";
import useAppState from "@/store/state";
import { LogIn, MoonIcon, SunIcon, UserIcon } from "lucide-react";
import { useNavigate } from "react-router-dom";

function Menubar() {
  const state = useAppState();
  const navigate = useNavigate();
  return (
    <MenubarPrimitive className="h-14" id="menubar" data-testid="menubar">
      <div className="flex-grow h-auto" id="menubar-spacer" />

      <MenubarMenu>
        <>
          {/* Register and Login buttons */}
          <Button
            className="my-1 px-1 py-1 h-12 flex flex-row text-base [&_svg]:size-6"
            variant={"outline"}
            onClick={() => navigate("/register")}
            data-testid="register-button"
          >
            <UserIcon data-testid="register-icon" />
            <p data-testid="register-text">Register</p>
          </Button>
          <Button
            className="my-1 px-1 py-1 h-12 flex flex-row text-base [&_svg]:size-6"
            variant={"outline"}
            onClick={() => navigate("/login")}
            data-testid="login-button"
          >
            <LogIn data-testid="login-icon" />
            <p data-testid="login-text">Login</p>
          </Button>
        </>
      </MenubarMenu>

      <MenubarMenu>
        <Button
          className="my-1 px-1 py-1 w-12 h-12 flex flex-row text-base [&_svg]:size-6"
          variant={"outline"}
          onClick={state.toggleMode}
          data-testid="mode-button"
        >
          {state.mode === "light" ? <MoonIcon data-testid="mode-icon" /> : <SunIcon data-testid="mode-icon" />}
        </Button>
      </MenubarMenu>
    </MenubarPrimitive>
  );
}

export default Menubar;
