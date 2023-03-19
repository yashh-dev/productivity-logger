import { applyMiddleware, combineReducers, createStore } from 'redux';
import thunk from 'redux-thunk';
import { persistReducer, persistStore } from 'redux-persist';
import storage from 'redux-persist/lib/storage';

import {DashboardMenuReducer } from './Reducer';

const persistConfig = {
    key: "productivity_logger_v1",
    whitelist: ["menu"],
    version: 1,
    storage: storage,
};

const reducers = combineReducers({
    menu: DashboardMenuReducer,
});

const persistedReducer = persistReducer(persistConfig, reducers);

const store = createStore(persistedReducer, applyMiddleware(thunk));

export const persistor = persistStore(store);

export default store;
