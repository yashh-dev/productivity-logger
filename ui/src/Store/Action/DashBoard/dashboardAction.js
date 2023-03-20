import { SET_MENU, REMOVE_MENU } from "../../Utils/actionType";

export const setMenu = (data) => ({
    type: SET_MENU,
    payload: data
});

export const removeMenu = () => ({
    type: REMOVE_MENU,
    payload: null
});