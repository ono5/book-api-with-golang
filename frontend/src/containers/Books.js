import React, { Component } from 'react';
import Book from '../components/Book';
import { books } from '../data';
import { connect } from 'react-redux';

class Books extends Component {
    constructor(props) {
        super(props);
    }

    render() {
        return (
            <div>
                <table className="table table-striped">
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Title</th>
                            <th>Author</th>
                            <th>Year</th>
                            <th>Actions</th>
                        </tr>
                    </thead>
                    <tbody>
                        {
                            this.props.books.map(book => {
                                return (
                                    <Book key={book.id} book={book}/>
                                )
                            })
                        }
                    </tbody>
                </table>
            </div>
        )
    }
}

const mapStateToProps = (state) => {
    return {
            books: state.booksData.books || [],
    };
};

export default connect(mapStateToProps, null)(Books);
