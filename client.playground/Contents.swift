import Foundation

enum ToDo {
    struct Request: Encodable {
        let title: String
        let content: String
    }
    
    struct Response: Codable {
        let id: String
        let title: String
        let content: String
        let due: Date
    }
}

let session: URLSession = .shared

let baseURL: URL = URL(string: "http://")!
