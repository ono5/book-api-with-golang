import {
    ADD_BOOK_ERROR,
    ADD_BOOK_LOADING,
    ADD_BOOK_SUCCESS,
    DELETE_BOOK_ERROR,
    DELETE_BOOK_LOADING,
    DELETE_BOOK_SUCCESS,
    EDIT_BOOK_ERROR,
    EDIT_BOOK_LOADING,
    EDIT_BOOK_SUCCESS,
    FETCH_BOOKS_ERROR,
    FETCH_BOOKS_LOADING,
    FETCH_BOOKS_SUCCESS
} from './types'

import axios from 'axios';

const url = 'http://localhost:8000/books';

export const fetchBooksSuccess = (data) => {
    return {
        type: FETCH_BOOKS_SUCCESS,
        payload: data,
    }
}

const normalizeResponse = (data) => {
    const arr = data.map(item => {
        const keys = Object.keys(item);

        keys.forEach(k => {
            item[k.toLowerCase()] = item[k];
            delete item[k];
        });

        return item;
    })

    return arr;
}

export const fetchBooks = () => {
    return (dispatch) => {
        return axios.get(url)
            .then(response => {
                const data = normalizeResponse(response.data);
                debugger
                dispatch(fetchBooksSuccess(data));
            }).catch(error => {

            });
    }
}
