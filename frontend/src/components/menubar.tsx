import { Button } from "@/components/ui/button";
import {
  Menubar as MenubarPrimitive,
  MenubarMenu,
  MenubarTrigger,
  MenubarContent,
  MenubarItem,
} from "@/components/ui/menubar";
import useAppState from "@/store/state";
import { Avatar, AvatarFallback, AvatarImage } from "@radix-ui/react-avatar";
import { LogIn, MoonIcon, SunIcon, UserIcon } from "lucide-react";
import { useEffect, useRef, useState } from "react";
import { useNavigate } from "react-router-dom";

function Menubar() {
  const intervalIdRef = useRef<NodeJS.Timeout | null>(null);
  const state = useAppState();
  const navigate = useNavigate();
  const [avatar, setAvatar] = useState<string | null>(null);

  useEffect(() => {
    const fetchAvatar = async () => {
      try {
        const response = await fetch(`${import.meta.env.VITE_API_URL ?? ""}/avatar`, {
          credentials: "include",
        });
        const data = await response.text();

        if (data === "NO") {
          return;
        }

        if (/^avatars\/.*\.png$/.test(data)) {
          setAvatar(data);
          if (intervalIdRef.current != null) {
            clearInterval(intervalIdRef.current);
          }
          intervalIdRef.current = null;
        }
      } catch (err) {
        console.error(err);
      }
    };

    if (avatar === null) {
      fetchAvatar();
      if (intervalIdRef.current != null) {
        intervalIdRef.current = setInterval(fetchAvatar, 100);
      }
    }

    return () => {
      if (intervalIdRef.current != null) {
        clearInterval(intervalIdRef.current);
      }
    };
  }, [avatar]);

  return (
    <MenubarPrimitive className="h-14" id="menubar" data-testid="menubar">
      {/* HomeHub apps */}
      <MenubarMenu>
        <MenubarTrigger>
          <strong>HomeHub</strong>
        </MenubarTrigger>
        <MenubarContent>
          <MenubarItem onClick={() => navigate("/pastebin")}>
            <strong>Pastebin</strong>
          </MenubarItem>
        </MenubarContent>
      </MenubarMenu>

      <div className="flex-grow h-auto" id="menubar-spacer" />

      {avatar != null ? (
        <MenubarMenu>
          <MenubarTrigger>
            <Avatar>
              <AvatarImage className="h-12 w-12 rounded-md" src={`${window.location.origin}/${avatar}`} />
              <AvatarFallback>
                <UserIcon />
              </AvatarFallback>
            </Avatar>
          </MenubarTrigger>
          <MenubarContent>
            <MenubarItem
              onClick={() => {
                fetch("/auth/logout", {
                  method: "POST",
                  credentials: "include",
                })
                  .then(() => navigate("/"))
                  .catch((err) => console.error(err));
              }}
            >
              <strong>Logout</strong>
            </MenubarItem>
          </MenubarContent>
        </MenubarMenu>
      ) : (
        <MenubarMenu>
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
        </MenubarMenu>
      )}

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
