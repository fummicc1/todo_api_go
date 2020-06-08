import Foundation

enum ToDo {
    struct Request: Encodable {
        let title: String
        let content: String?
    }
    
    struct Response: Decodable {
        let id: String
        var title: String
        var content: String?
        var due: Date?
    }
}

let session: URLSession = .shared

let baseURL: URL = URL(string: "http://localhost:8080/api/v1/todo")!

func get() {
    var request: URLRequest = URLRequest(url: baseURL)
    request.httpMethod = "GET"
    session.dataTask(with: request) { (data, response, error) in
        if let error = error {
            fatalError(error.localizedDescription)
        }
        guard let response = response as? HTTPURLResponse, response.statusCode == 200 else {
            fatalError()
        }
        guard let data = data else {
            fatalError()
        }
        print(data)
        guard let todos = try? JSONDecoder().decode([ToDo.Response].self, from: data) else {
            fatalError()
        }
        print(todos)
    }.resume()
}


func post() {
    var request: URLRequest = URLRequest(url: baseURL)
    request.httpMethod = "POST"
    let todo: ToDo.Request = .init(title: "AAA", content: "BBBB")
    request.httpBody = try? JSONEncoder().encode(todo)
    print(request.httpBody)
    session.dataTask(with: request) { (data, response, error) in
        if let error = error {
            fatalError(error.localizedDescription)
        }
        print(response)
        guard let response = response as? HTTPURLResponse, response.statusCode == 200 else {
            fatalError()
        }
        print(response)
    }.resume()
}

post()
