import React, { Component } from 'react'
import {Button, Card, Header, Form, Input, Icon, TextArea } from "semantic-ui-react";
import axios from 'axios'
class PostForm extends Component {
	constructor(props) {
		super(props)

		this.state = {
			id: '',
			title: '',
			text: '',
			edit: false,
			editId: '',
		}
	}
	componentDidMount() {
		this.getNote();
	}
	changeHandler = e => {
		this.setState({ [e.target.name]: e.target.value })
	}
    onEdit(idEd) {
		this.setState({ edit: true,editId: idEd})
		this.getNote()
	} 
	getNote = () => {
    axios.get("/api/retrieve").then(res => {
      if (res && res.data) {
        this.setState({
          items: res.data.map(item => {

            return (
				
              <Card id={item.id}  fluid>
				  
				{ this.state.editId === item.id ? (
						<Form >
							<Form.Field>
								<Input
									type="text"
									name="title"
									defaultValue={item.title}
									onChange={this.changeHandler}
									fluid
									placeholder="Title of note"
									required
								/>
							</Form.Field>
							<Form.Field>
								<TextArea
									type="text"
									name="text"
									defaultValue={item.text}
									onChange={this.changeHandler}
									fluid
									placeholder="Content of note"
								/>
							</Form.Field>
							<Button onClick={() => this.updateHandler(item.id)}>Edit Note</Button>
						</Form>

				):(
				  
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={{ wordWrap: "break-word" }}>{item.id} {item.title}</div>
                    <br/>
                    <div style={{ wordWrap: "break-word" }}>{item.text}</div>
                  </Card.Header>

                  <Card.Meta textAlign="right">
                    <Button
                      name="edit"
                      color="yellow"
					  onClick={() => this.onEdit(item.id)}
                    />
                    <span style={{ paddingRight: 10 }}>Edit</span>
                    <Button
                      name="delete"
                      color="red"
					  onClick={() => this.deleteHandler(item.id)}
                    />
                    <span style={{ paddingRight: 10 }}>Delete</span>
                  </Card.Meta>
                </Card.Content>
				)}
              </Card>
            );
          })
        });
      } else {
        this.setState({
          items: []
        });
      }
    });
  };
	submitHandler = e => {
		e.preventDefault()
		console.log(this.state)
		axios
			.post('/api/store', this.state)
			.then(response => {
				console.log(response)
			})
			.catch(error => {
				console.log(error)
			})
		window.location.reload(false);
	}
	deleteHandler(id) {
	//	e.preventDefault()
		const url = '/api/delete/'+id
		console.log(this.state)
		axios
			.delete(url)
			.then(response => {
				console.log(response)
			})
			.catch(error => {
				console.log(error)
			})
		window.location.reload(false);
	}
	editHandle(id){
		this.setState({ edit: true ,editId: id})
	}
	updateHandler(id) {
	//	e.preventDefault()
		const url = '/api/update/'+id
		console.log(this.state)
		axios
			.put(url, this.state)
			.then(response => {
				console.log(response)
			})
			.catch(error => {
				console.log(error)
			})
		window.location.reload(false);
		}


	render() {
		const { id, title, text} = this.state
		const show=!this.state.edit
		return (
			<div>
				<div className="row">
				<Header as='h1' content='MyNote APP' style={{marginTop: '3em',marginBottom: '3em'}} textAlign='center' />
				</div>
					<div className="row">{show &&
						<Form onSubmit={this.submitHandler}>
							<Form.Field>
								<Input
									type="text"
									name="title"
									value={title}
									onChange={this.changeHandler}
									fluid
									placeholder="Title of note"
									required
								/>
							</Form.Field>
							<Form.Field>
								<TextArea
									type="text"
									name="text"
									value={text}
									onChange={this.changeHandler}
									fluid
									placeholder="Content of note"
								/>
							</Form.Field>
							<Button type="submit">Create Note</Button>
						</Form>}
					</div>
					<br></br>
					<div className="row">
				<Card.Group>{this.state.items}</Card.Group>
				</div>
			</div>
		)
	}
}

export default PostForm
