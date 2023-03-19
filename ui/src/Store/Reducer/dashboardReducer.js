import { SET_MENU, REMOVE_MENU} from "../Utils/actionType";

const DashboardMenuReducer = (state=null, action) => {
    switch (action.type) {
        case SET_MENU:
            return {menu: action.payload};
        case REMOVE_MENU:
            return null;
        default:
            return state;
    }
}

export default DashboardMenuReducer;