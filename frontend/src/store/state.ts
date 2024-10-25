import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

type State = {
  mode: "light" | "dark";
};

interface StateAction {
  toggleMode: () => void;
}

const useAppState = create<State & StateAction>()(
  persist(
    (set) => ({
      mode: "light",
      toggleMode: () =>
        set((state) => {
          if (state.mode === "light") {
            document.documentElement.classList.remove("dark");
            document.documentElement.classList.add("light");
          } else {
            document.documentElement.classList.remove("light");
            document.documentElement.classList.add("dark");
          }

          return { ...state, mode: state.mode === "light" ? "dark" : "light" };
        }),
    }),
    {
      name: "state",
      storage: createJSONStorage(() => localStorage),
    }
  )
);

export default useAppState;
