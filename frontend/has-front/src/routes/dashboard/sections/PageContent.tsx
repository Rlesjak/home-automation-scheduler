
interface t_PageContentProps {
    className?: string
    children?: React.ReactNode
}

function PageContent(props: t_PageContentProps) {

    
    return (
        <div className={(props.className || "") + ""}>
            { props.children || null }
        </div>
    )
}

export default PageContent