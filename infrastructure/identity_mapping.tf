data "aws_iam_policy_document" "assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["pods.eks.amazonaws.com"]
    }

    actions = [
      "sts:AssumeRole",
      "sts:TagSession"
    ]
  }
}

resource "aws_iam_role" "eks-pod-identity" {
  name               = "eks-bedrock-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy_attachment" "attach_bedrock" {
  policy_arn = "arn:aws:iam::aws:policy/AmazonBedrockFullAccess"
  role       = aws_iam_role.eks-pod-identity.name
}

resource "aws_eks_pod_identity_association" "eks_service_account" {
  cluster_name    = "eks-cluster"
  namespace       = "default"
  service_account = "eks-service-account"
  role_arn        = aws_iam_role.eks-pod-identity.arn
  depends_on = [ module.eks ]
}
