import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

export type State = {
  mode: "light" | "dark";
};

export interface StateAction {
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
          } else {
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
