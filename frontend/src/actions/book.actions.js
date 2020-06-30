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
import { history } from '../index';

const url = 'http://localhost:8000/books';

// CREATE----------------------------------------
export const createBookSuccess = (data) => {
    return {
        type: ADD_BOOK_SUCCESS,
        payload: data,
    }
}

export const createBook = (book) => {
    const data = {
        title: book.title,
        author: book.author,
        year: book.year,
    }

    return (dispatch) => {
        return axios.post(url, data)
            .then(response => {
                const id = response.data;
                axios.get(`${url}/${id}`)
                    .then(response => {
                        const data = response.data;
                        const normalizeData = {
                            id: data.id,
                            title: data.title,
                            author: data.author,
                            year: data.year,
                        };

                        dispatch(createBookSuccess(normalizeData));
                        history.puah('/');
                    }).catch(error => {
                        console.log(error);
                    })
            }).catch(error => {
                console.log(error);
            });
    }
}

// EDIT------------------------------------------

// DELETE----------------------------------------

// FETCH-----------------------------------------

export const fetchBooksSuccess = (data) => {
    return {
        type: FETCH_BOOKS_SUCCESS,
        payload: data,
    }
}

export const fetchBooksLoading = (data) => {
    return {
        type: FETCH_BOOKS_LOADING,
        payload: data,
   };
}

export const fetchBooksError = (data) => {
    return {
        type: FETCH_BOOKS_ERROR,
        payload: data,
    };
}

export const fetchBooks = () => {
    let isLoading = true;

    return (dispatch) => {
        dispatch(fetchBooksLoading(isLoading));
        return axios.get(url)
            .then(response => {
                const data = response.data;
                dispatch(fetchBooksSuccess(data));
                isLoading = false;
                dispatch(fetchBooksLoading(isLoading));
            }).catch(error => {
                const errorPayload = {};
                errorPayload['message'] = error.response.statusText;
                errorPayload['status'] = error.response.status;
                dispatch(fetchBooksError(errorPayload));
                isLoading = false;
                dispatch(fetchBooksLoading(isLoading));
            });
    }
}
